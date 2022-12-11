resource "digitalocean_app" "timer-function" {
  spec {
   name = "timer-function" 
   region = "ams"
   domain {
      name = "timer-relay.cooperkyle.com" 
   }
   function {
      name = "relay-switcher-timer${local.rpi_name}-${local.rpi_id}" 
      github {
        repo = "kc8/relay-switcher"
        branch = "timer-function"
        deploy_on_push = true
      }
   }
  } 
}
