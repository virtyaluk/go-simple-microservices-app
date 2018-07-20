package controllers

import (
  "gopkg.in/mgo.v2"
  "github.com/virtyaluk/go-simple-microservices-app/bookings/common"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
  MongoSession *mgo.Session
}

// Close mgo.Session
func (context *Context) Close() {
  context.MongoSession.Close()
}

// Return mgo.Collection for the given name
func (context *Context) DbCollection(name string) *mgo.Collection {
  return context.MongoSession.DB(common.AppConfig.Database).C(name)
}

// Create a new Context object for each HTTP request
func NewContext() *Context {
  session := common.GetSession().Copy()
  context := &Context{
    MongoSession: session,
  }

  return context
}
