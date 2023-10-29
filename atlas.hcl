variable "cloud_token" { 
  type    = string
  default = getenv("ATLAS_TOKEN")
}
  
atlas {
  cloud {
    token = var.cloud_token
  }
}

variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql+wss://datum-matoszz.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
  migration {
    dir = "file://db/migrations"
    format = atlas
  }
}