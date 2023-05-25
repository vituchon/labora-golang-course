# Disposición del código en el sistema de archivos / Code layout in file system

/─┬─model/ : The abstractions regarding the domain itself and its internal logic. Unit tests missing for now.
  │
  ├─repositories/ : Handling of model abstractions that exists over an underlying persistence storage
  │
  ├─presentation/─┐ : Code regarding exposing to clients a way to work with the model and "delivering" meaningful data
  │               ├─web/─┐ : HTTP REST API
  |               |      ├─ controllers/  : receving requests, process and dispaching responses (uses repositories to work with persistent data)
  │               │      └─ server.go     : ends points declaration (implemented on controllers/) and server fine tune
  │               │
  │               └─cli/ : A very basical implementation of the game using the command line (COMMING SOON!)
  │
  └─main.go: on launch.. code executes from here.
