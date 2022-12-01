# Resource = Bloco
# Local = Provider
# File = Provider Type

resource "local_file" "example" {
  content = "foo"
  filename = "example.txt" 
}

