workflow "Build JSON" {
  resolves = ["../action/"]
  on = "schedule(0 22 * * *)"
}

action "../action/" {
  uses = "./action"
  env = {
    PAGES_BRANCH = "master"
  }
  secrets = ["TOKEN"]
}
