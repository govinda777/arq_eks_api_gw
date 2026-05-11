variable "db_name" {
  description = "Database name"
  type        = string
  default     = "itemsapi"
}

variable "db_username" {
  description = "Database admin username"
  type        = string
  default     = "admin"
}

variable "db_password" {
  description = "Database admin password"
  type        = string
  sensitive   = true
}
