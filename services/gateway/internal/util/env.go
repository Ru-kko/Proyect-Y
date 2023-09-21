package util

import ( 
  ev "Proyect-Y/common/env"
)

type Enviroment struct {
  PORT int
  CONFIG string
}

var env *Enviroment


func GetEnv() Enviroment {
  if  env != nil {
    return *env
  }

  env = &Enviroment{}

  ev.Load()
  ev.Parse[Enviroment](env)

  return *env
}
