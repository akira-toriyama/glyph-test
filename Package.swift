// swift-tools-version:6.0
// Live-fire fixture (t-5sas): declares the family floor so the dead-gate
// check below it has something to compare against.
import PackageDescription

let package = Package(
    name: "Fixture",
    platforms: [.macOS("26.0")],
    targets: [.target(name: "Fixture")]
)
