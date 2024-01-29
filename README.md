# Evaluation OPS DIIAGE 3


Le but de cette evaluation est de valider vos competences sur les concepts de securite, les concepts avances (operateurs) et le monitoring d'application. Prenez le temps de lire tout l'ennonce avant de vous lancer.

**Tout votre travail (fichiers de config, YAML, reponses aux questions additionelles, etc) sont a mettre dans le dossier `/reponses`. A la fin du temps imparti, veuillez faire un zip de ce dossier et l'envoyer a Luc CHMIELOWSKI sur Teams.**

**:warning: En cas de soucis avec l'enonce n'hesitez pas a contacter Luc CHMIELOWSKI directement**


## Contexte

Nous avons developpe une application (un blog) et sommes pret a la deployer en production, mais avant de la rendre publique, nous voulons creer tout le monitoring necessaire pour s'assurer qu'elle tourne parfaitement.

### Setup: Creation du cluster kind et deploiement de l'application test

Dans cet exercice le code de deploiement de l'application et du cluster sont deja fournis.

Afin de creer le cluster kind vous pouvez utiliser la commande suivante:

```bash
kind create cluster --config=kind.yaml
```

Une fois le cluster correctement lance, vous pouvez deployer notre blog en lancant:

```bash
kubectl apply -f application.yaml
# Attendre que le pod de l'application soit lance
kubectl wait --for=condition=Ready pod -l app=blog -n blog --timeout=4m
```

Afin de verifier que le code tourne correctement, vous pouvez utiliser la commande suivante (Notre service utilse un service de type `NodePort` pour exposer notre blog sur le port 8080 de notre machine, on peut donc l'appeler directement) :

```bash
curl localhost:8080/api/healthz | jq
```

Vous devriez avoir la reponse suivante:

```json
{
  "alive": true
}
```

En allant sur votre navigateur au `http://localhost:8080/` vous verrez l'application. Cette derniere compte le nombre de hits sur le endpoint `/`

## Exercice principal

L'application en question emet des metrics qui peuvent etre scrappees par Prometheus sur son endpoint `/api/metrics`.

Entant qu'ops nous voulons donc :
- installer `prometheus-operator`
- configurer une instance prometheus pour recuperer les metrics emises par notre blog
- Mettre en place 2 alertes:
  - la premiere qui nous `page` quand on a atteint plus de 20 visite sur notre site
  - la seconde qui nous `page` quand on le pod du blog est dans un mauvais etat pour une periode > 10 secondes

## Questions additonelles

Une fois l'exercice ci dessus termine, vous pouvez repondre a ces questions:

1. Actuellement le compteur de visite est local au pod du blog, proposez un ou plusieurs changements pour que notre service et les alertes mises en place continuent de fonctionner correctement si on decide d'ajouter des pods.

2. Par definition un `controlleur` a besoin de lire et ecrire des resources dans kubernetes. D'un point de vue securite, comment donne t'on les acces necessaires a notre pod ? Donnez un exemple d'implementation permettant aux pod(s) de l'operateur prometheus de `get` toutes les resources dans le namespace `demo`

3. Expliquez en quelques mots le fonctionnement du `prometheus-operator` (uniquement de l'operateur, on ne parlera pas ici de prometheus en lui meme).

4. Nous verifions actuellement que notre service ne crash pas. Devrions nous verifier d'autres indicateurs ?

5. Un CRD vient souvent de pair avec des `admission webhook` que font ils et a quoi servent ils ? Quelle est la difference entre un `validation webhook` et un `mutation webhook` ?

## Liens utiles

- [fichier YAML d'installation de prometheus-operator](https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml) (`kubectl create -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/main/bundle.yaml`)
- [documentation](https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/getting-started.md)

## Notes:


- **Vous n'avez normalement pas besoin de Grafana, un serveur rometheus possede une GUI sur le port `9090` par defaut**
- Une fois l'interface de prometheus atteinte, vous pouvez faire la query suivante : `{container="blog"}` afin de lister toutes les metrics pour notre container
- Si vous rencontrez l'erreur suivante :`The CustomResourceDefinition "prometheuses.monitoring.coreos.com" is invalid: metadata.annotations: Too long: must have at most 262144 bytes`, utilisez `kubectl create -f` au lieu de `kubectl apply -f`
- Si vous avez des problemes pour charger l'image du blog depuis dockerhub, vous pouvez la construire et l'importer dans `kind` en faisant:

```
docker build -t luskidotme/blog:v1 application/
kind load docker-image -n diiage luskidotme/blog:v1
```