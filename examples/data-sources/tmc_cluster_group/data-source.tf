# Read Tanzu Mission Control cluster group : fetch cluster group details
data "tmc_cluster_group" "read_cluster_group" {
  name = "default"
}