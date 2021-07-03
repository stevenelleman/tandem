## Ingress Rules

The Ingress resource is a rule set used for the Ingress Controller to route traffic. 

Without the Ingress Controller the rule cannot be applied. Installation can be found [here](https://kubernetes.github.io/ingress-nginx/deploy/#aws).

There are two ingress resources because _by default_ the ingress controller removed the matching prefix.

Therefore traffic for `/v1/silos` was rewritten to be `/silos`, which is incorrect. 

To retain our current public api prefix/routes a second ingress resource was added with a rewrite rule to include the prefix. 