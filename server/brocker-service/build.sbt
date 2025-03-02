import Dependencies._

ThisBuild / scalaVersion     := "2.13.12"
ThisBuild / version          := "0.1.0-SNAPSHOT"
ThisBuild / organization     := "com.example"
ThisBuild / organizationName := "example"

lazy val root = (project in file("."))
  .settings(
    name := "brocker-service",
    libraryDependencies += munit % Test
  )

libraryDependencies += "org.postgresql" % "postgresql" % "42.5.4"
libraryDependencies += "com.fasterxml.jackson.module" %% "jackson-module-scala" % "2.14.0"
// See https://www.scala-sbt.org/1.x/docs/Using-Sonatype.html for instructions on how to publish to Sonatype.
