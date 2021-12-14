## fibonacci service, non performant

Avec une complexité de O(n^2), ne pas demander la valeur au dessus de 45, le service galèrera sa race :)

---

## set autoscaler

kubectl autoscale deployment --namespace hpa horizontal-scaling --cpu-percent=30 --min=2 --max=5

---

## suite dans liveness

---