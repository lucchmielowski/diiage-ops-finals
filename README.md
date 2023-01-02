# Evaluation OPS DIIAGE 3 P1 & 2

L'evaluation est une evaluation pratique visant a evaluer votre capacite a admnistrer un cluster kubernetes

# Prerequis

- Lancez la commande `docker-compose up -d` lancer les 3 noeuds de votre cluster kubernetes (ces derniers sont pour l'instant vide a l'exception de tous les outils necessaire pour installer un cluster et ne forment pas encore un cluster kubernetes)

- Vous pouvez ensuite vous connecter sur les noeuds grace a la commande `docker exec -it <nom du noeud> bash`

- Vous pouvez lister les noeuds depuis votre machine en faisant un `docker ps`

- Tout le contenu du dossier `./deployments` est monte dans le noeud `k8s-control-plane-1` au path `/root/app`

# Exercice

### Question 1
En utilisant `kubeadm` boostrapez les noeud pour creer un cluster avec la topologie suivante :
  - `k8s-control-plane-1`: master
  - `k8s-worker-1` et `k8s-worker2-1`: nodes

A l'issue de cette question, la commande `kubectl get nodes` devra afficher des noeuds comme etant `Ready`.

### Question 2

Via la commande `helm` et la [chart de bitnami](https://artifacthub.io/packages/helm/bitnami/redis) installez un redis:
 - dans le namespace `catalog`
 - avec les values contenues dans le fichier dans `deployment/redis-values.yaml` (`/root/app/redis-values.yaml` sur le noeud control-plane)

Pour installer helm:
```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

Assurez vous que le redis tourne correctement

Appliquez ensuite le fichier `catalog.yaml` dans le cluster

Assurez vous que les pods `catalog` fonctionnent correctement

> HINT rancher/local-path-provisioner


### Question 3

Appliques le fichier `deployment/presentation.yaml` (`/root/app/presentation.yaml`) dans le namespace `presentation`

Assurez vous ensuite que vous etes capables d'appeler le service depuis _votre machine_ en faisant un `curl localhost:8080/`

Vous devriez recevoir un `HTTP 200` et une reponse au format texte

> HINT: Voici la configuration du cluster `kind` qu'on avait pour travailler en local:
  ```
  kind: Cluster
  apiVersion: kind.x-k8s.io/v1alpha4
  name: diiage
  nodes:
      - role: control-plane
      kubeadmConfigPatches:
        - |
          kind: InitConfiguration
          nodeRegistration:
            kubeletExtraArgs:
              node-labels: "ingress-ready=true"
      extraPortMappings:
        - containerPort: 80
          hostPort: 8080
          protocol: TCP
        - containerPort: 443
          hostPort: 8081
          protocol: TCP
    - role: worker
  ```
> A noter que vous pouvez ajouter de la configuration sur les noeuds deja existants via `kubectl`

### Question 4

Modifiez et redeployez le fichier `deployments/catalog.yaml` afin:
- D'avoir 3 replicas
- De n'autoriser le traffic du service a arriver uniquement qu'a partir du moment ou le endpoint `/get-key` peut traiter les requetes GET
- Un redemarrage automatique du pod si le endpiont `/get-key` ne repond plus aux requetes GET
- Reserver 100Mi de Memoire et 100m de CPU aupres du scheduler de kubernetes
