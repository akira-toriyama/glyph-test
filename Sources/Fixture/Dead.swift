// Live-fire fixture: this gate sits BELOW the declared floor (26.0) and must
// be flagged as dead code.
func effect() {
    if #available(macOS 14, *) {
        modern()
    } else {
        legacyFallback()
    }
}
func modern() {}
func legacyFallback() {}
