# Resource = Bloco
# Local = Provider
# File = Provider Type

resource "local_file" "example" {
  content = var.content
  filename = "example.txt" 
}

variable "content" {
  type = string
  default = "hello world"
}