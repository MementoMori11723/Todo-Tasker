export function load() {
  return fetch("/api")
    .then((res) => res.json())
    .catch((err) => console.error(err.message));
}
