// Live-fire fixture, pass side: this gate sits ABOVE the declared floor
// (26.0) — a legitimate new-API adoption gate the check must let through.
func effect() {
    if #available(macOS 27, *) {
        adoptNewAPI()
    }
}
func adoptNewAPI() {}
