resource "loadmaster_virtual_service" "foo" {
  address  = "192.168.1.10"
  protocol = "tcp"
  port     = "8080"
  nickname = "bar"
  type     = "gen"
}


resource "loadmaster_real_server" "baz" {
  virtual_service_id = loadmaster_virtual_serice.foo.id
  address            = "192.168.1.11"
  port               = "8080"
}
