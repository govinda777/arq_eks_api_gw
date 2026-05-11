
# ITEMS API - PRODUCTION MODEL #

## 1 - FRAGMENTAÇÃO DE TRÁFEGO

Esta aplicação foi desenvolvida em Golang seguindo as melhores práticas de arquitetura e implantação em nuvem, servindo como um modelo de referência para APIs de alto desempenho.

**1.1 - API ITEMS**

- Construir uma aplicação em Golang que retorna uma lista de itens de forma eficiente.
- Implementação de Logs estruturados
- Integração com RDS (AWS)
- Dockerização otimizada
- Manifestos Kubernetes para orquestração

**1.2 - Provisionamento Cluster EKS**
- Infraestrutura como Código (Terraform)
- Utilização de instâncias Spot para otimização de custos
- Instalação de Add-Ons essenciais (Kube-Proxy, CoreDNS, VPC CNI, EBS CSI Driver)
- Observabilidade de custos com Kubecost
- Balanceamento de carga com AWS Application Load Balancer (ALB)
- Regras de roteamento por Header para múltiplos ambientes (MLA/MLB/MLC/MLM)

**1.3 - Persistência de Dados**
- Instância gerenciada RDS MySQL

> Para fins de demonstração foi utilizado um RDS compartilhado, permitindo a segregação lógica por tabelas ou schemas conforme a necessidade de performance e isolamento.

O roteamento de tráfego é realizado de forma inteligente pelo Application Load Balancer com base nos Headers da requisição, eliminando a necessidade de proxies complexos e reduzindo a latência.

Para representar a solução foi desenhado o diagrama de arquitetura:

![DIAGRAMA DE ARQUITETURA ](https://github.com/marcosouzatech/desafio/blob/main/img/diagrama.png)


Exemplo de requisição para o ambiente da API Items:
```
curl -H "site_id":"MLB" https://api.prod.example.com/api/items
```

## 2 - LIMITAÇÃO DE TRÁFEGO

A API conta com mecanismos de Rate Limiting para garantir a disponibilidade e proteção contra abusos.

A configuração atual limita o tráfego em 15 requisições por segundo, garantindo conformidade com o limite de 1000 requisições por minuto.

Ao atingir o limite, a API retorna o status `429 TOO MANY REQUESTS`.

![TESTE DE CARGA POSTMAN](https://github.com/marcosouzatech/desafio/blob/main/img/teste_postman.png)

![TESTE DE CARGA HEY](https://github.com/marcosouzatech/desafio/blob/main/img/teste_hey.png)

## 3 - MONITORAMENTO E OBSERVABILIDADE

Utilizamos o NewRelic para monitoramento Full-Stack, incluindo:
- Métricas de Infraestrutura (EKS/Nodes)
- Tracing Distribuído
- Agregação de Logs
- Telemetria com OpenTelemetry

Dashboards personalizados permitem visualizar em tempo real o Error Rate, Latência (p99) e consumo de recursos.

[DASHBOARD EXEMPLO - PDF](https://github.com/marcosouzatech/desafio/blob/main/img/dashboard_argentina.pdf)

## CUSTOS

A gestão de custos é realizada via Kubecost, fornecendo visibilidade granular sobre os gastos do cluster Kubernetes.

![KUBECOST OVERVIEW](https://github.com/marcosouzatech/desafio/blob/main/img/kubecost.png)

## CI/CD - GITHUB ACTIONS

Pipeline automatizada para Build e Publish da imagem no DockerHub. O processo de deploy é integrado via GitOps.

## Como Rodar o Projeto Localmente

### Pré-requisitos
- Go 1.20 ou superior
- MySQL
- Docker (opcional)

### Configuração do Ambiente
1. Clone o repositório.
2. Navegue até a pasta `items-api/`.
3. Copie o arquivo de exemplo de variáveis de ambiente:
   ```bash
   cp .env.example .env
   ```
4. Edite o arquivo `.env` com as suas credenciais do banco de dados.

### Execução
#### Via Go Local:
```bash
cd items-api/
go run main.go
```

#### Via Docker:
```bash
docker build -t items-api .
docker run -p 9000:9000 --env-file items-api/.env items-api
```

### Variáveis de Ambiente
| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `API_PORT` | Porta em que a API irá escutar | `9000` |
| `DB_USUARIO` | Usuário do banco de dados | - |
| `DB_SENHA` | Senha do banco de dados | - |
| `DB_URL` | Host do banco de dados | `localhost` |
| `DB_PORT` | Porta do banco de dados | `3306` |
| `DB_NOME` | Nome do banco de dados/schema | `brasil` |
| `APP_NAME` | Nome da aplicação no NewRelic | `api-items` |
| `NEW_RELIC_LICENSE_KEY` | Chave de licença do NewRelic | - |

##### CONTATOS
- Linkedin: (https://www.linkedin.com/in/marcosouzatech/)
- Youtube: https://www.youtube.com/@maistalkmenosshow
- Email: marcos.souza@luvtech.com.br
