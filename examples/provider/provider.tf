provider "loadmaster" {
  server  = "192.168.1.10"
  api_key = "<SECRET API KEY>"

  # 1 = XML 2 = JSON (firmware version >=7.2.50)
  api_version = 2
}
