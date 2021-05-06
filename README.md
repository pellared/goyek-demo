# goyek-demo

> Demonstration of [goyek](https://github.com/goyek/goyek)

https://www.youtube.com/watch?v=f-r_9IJ0xJc&t=558s

[![Build Status](https://img.shields.io/github/workflow/status/pellared/goyek-demo/build)](https://github.com/pellared/goyek-demo/actions?query=workflow%3Abuild+branch%3Amain)

[![Demo](https://img.youtube.com/vi/f-r_9IJ0xJc/hqdefault.jpg)](https://youtu.be/f-r_9IJ0xJc?t=558)

[Presentation](https://docs.google.com/presentation/d/1GUEM9oUUi1h23Tknws8hkNdjXQt74nyA1Saqv0IHRE0/edit?usp=sharing)

Running:

```sh
go run ./build -h
```

Notable files:

- [build/build.go](build/build.go) - taskflow build pipeline,
- [.github/workflows/build.yml](.github/workflows/build.yml) - GitHub Actions workflow,
- [.vscode/launch.json](.vscode/launch.json) - Visual Studio Code configuration for debugging,
- [.vscode/tasks.json](.vscode/tasks.json) - Visual Studio Code task definition for runing taskflow using `Tasks: Run Build Task`.
