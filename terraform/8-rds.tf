resource "aws_db_instance" "items-api-db" {
  
  allocated_storage       = 10
  identifier              = "items-api"
  engine                  = "mysql"
  engine_version          = "5.7"
  instance_class          = "db.t3.micro"
  db_name                 = var.db_name
  username                = var.db_username
  password                = var.db_password
  parameter_group_name    = "default.mysql5.7"
  publicly_accessible     = false
  skip_final_snapshot     = true
}
 
output "Endpoint_string" {
  value = aws_db_instance.items-api-db.endpoint
 
}