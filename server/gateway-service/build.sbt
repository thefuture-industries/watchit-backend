import Dependencies._

ThisBuild / scalaVersion     := "2.13.12"
ThisBuild / version          := "0.1.0"
ThisBuild / organization     := "com.flicksfi"
ThisBuild / organizationName := "flicksfi"

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
libraryDependencies += "org.scalatest" %% "scalatest" % "3.2.16" % Test

enablePlugins(ScalafmtPlugin)
enablePlugins(AssemblyPlugin)

assembly / assemblyMergeStrategy := {
  case PathList("META-INF", xs @ _*) => MergeStrategy.discard
  case "reference.conf" => MergeStrategy.concat  // Объединяет все reference.conf
  case "application.conf" => MergeStrategy.concat  // Объединяет все application.conf
  case "version.conf" => MergeStrategy.first  // Берёт первую версию version.conf
  case _ => MergeStrategy.first
}

// See https://www.scala-sbt.org/1.x/docs/Using-Sonatype.html for instructions on how to publish to Sonatype.
