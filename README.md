# DevoxxFR 2025

![](img/logo.png)

- Titre : [**Infisical : Le meilleur ami des devs, pour des secrets bien gardés !**](https://mobile.devoxx.com/events/devoxxfr2025/talks/5874/details)
- Speaker : **Julien Briault**

- Abstract :

  - La gestion des secrets est un enjeu crucial pour la sécurité des infrastructures modernes. Si HashiCorp Vault a longtemps  régné en maître, de nouvelles alternatives Open Source viennent aujourd'hui en proposant des alternatives sérieuses.
  - Je présenterais Infisical, une solution Open Source qui mise sur la simplicité d’utilisation, l’efficacité, et une compatibilité sans faille avec vos environnements de développement préférés.
  - Spoiler : gérer vos secrets n’a jamais été aussi simple, même avec des environnements multiples (prod/staging/dev/sandbox) – de quoi enfin donner à vos développeurs un sommeil un peu plus paisible.
  - Et ce n’est pas tout : depuis qu’Hashicorp (IBM) a pris un virage vers la BSL (Business Source License), certains challengers comme Infisical qui redouble d'arguments pour vous séduire.
  - Enfin, pour ne pas vous laisser sur votre faim, je conclurais avec une démonstration technique. Ainsi, vous découvrirez comment intégrer Infisical dans Kubernetes (grâce au Secret Operator dédié), mais aussi avec Ansible et plusieurs langages comme Java, Python et Go.
  - Je vous le promets, après cette conférence, gestion des secrets n'aura plus aucun secret pour vous !

## La structure du projet

```bash
.
├── README.md # Vous êtes ici !
├── demo
│   ├── README.md # Les anti-sèches, parce que, pourquoi pas ?
│   ├── go_get_secret # La démo
│   │   ├── devoxx
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── handlers.go
│   │   ├── main.go
│   │   ├── templates
│   │   │   └── logs.html
│   │   └── utils.go
│   ├── project_example # Exemple de projet
│   └── project_with_secrets # Projet avec des secrets à scanner
│       ├── go.mod
│       └── main.go
├── img # De belles images
│   └── logo.png
└── slides # Les slides du talk
```

## Les démonstrations

Les démos sont réutilisables dans le dossier `demo/`.
