import Dependencies._

ThisBuild / scalaVersion     := "2.13.12"
ThisBuild / version          := "0.1.0"
ThisBuild / organization     := "com.flicksfi"
ThisBuild / organizationName := "flicksfi"

lazy val root = (project in file("."))
  .settings(
    name := "broker-service",
    libraryDependencies += munit % Test
  )

javacOptions ++= Seq("-encoding", "UTF-8")
scalacOptions ++= Seq("-encoding", "UTF-8")

libraryDependencies += "org.postgresql" % "postgresql" % "42.5.4"
libraryDependencies += "com.fasterxml.jackson.module" %% "jackson-module-scala" % "2.14.0"

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
