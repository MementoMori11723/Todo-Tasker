import type { PageLoad } from "./$types";
export const ssr = false;
export const load = (async (loadEvent) => {
  try {
    const { fetch } = loadEvent;
    const res = await fetch("/api");
    const data = await res.json();
    return data;
  } catch (e: any) {
    return { error: e.message };
  }
}) satisfies PageLoad;
