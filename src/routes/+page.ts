import type { PageLoad } from "./$types";
export const ssr = false;
export const load = (async () => {
  const res = await fetch("/api");
  const data = await res.json();
  return data;
}) satisfies PageLoad;
