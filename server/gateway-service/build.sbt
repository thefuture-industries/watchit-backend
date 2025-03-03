import Dependencies._

ThisBuild / scalaVersion     := "2.13.12"
ThisBuild / version          := "0.1.0-SNAPSHOT"
ThisBuild / organization     := "com.example"
ThisBuild / organizationName := "example"

lazy val root = (project in file("."))
  .settings(
    name := "gateway-service",
    libraryDependencies += munit % Test
  )


libraryDependencies += "com.typesafe.akka" %% "akka-actor" % "2.6.20"
libraryDependencies += "com.typesafe.akka" %% "akka-stream" % "2.6.20"
libraryDependencies += "com.typesafe.akka" %% "akka-http" % "10.2.10"
libraryDependencies += "com.typesafe.akka" %% "akka-http-spray-json" % "10.2.10"
libraryDependencies += "ch.qos.logback" % "logback-classic" % "1.2.11"

// See https://www.scala-sbt.org/1.x/docs/Using-Sonatype.html for instructions on how to publish to Sonatype.
