set autoscaler:

kubectl autoscale deployment --namespace hpa horizontal-scaling --cpu-percent=30 --min=1 --max=5