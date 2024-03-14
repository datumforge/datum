variable "cloud_token" { 
  type    = string
  default = getenv("ATLAS_CLOUD_TOKEN")
}
  
atlas {
  cloud {
    token = var.cloud_token
  }
}

data "remote_dir" "migrations" {
  name = "datum"
}


variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql://datum-datum.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
  migration {
    dir = data.remote_dir.migrations.url
  }
}