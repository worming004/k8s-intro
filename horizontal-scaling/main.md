set autoscaler:

kubectl autoscale deployment --namespace hpa horizontal-scaling --cpu-percent=30 --min=2 --max=5
