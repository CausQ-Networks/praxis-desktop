import { useEffect, useRef } from "react";
import { Events } from "@wailsio/runtime";

/** Subscribes to a backend event for the lifetime of the component. */
export function useEvent<T>(name: string, handler: (data: T) => void) {
  const ref = useRef(handler);
  ref.current = handler;
  useEffect(() => Events.On(name, (ev) => ref.current(ev.data as T)), [name]);
}

/** Logs a failed backend call; calls never throw into the UI. */
export function report(p: Promise<unknown>) {
  p.catch((err) => console.error(err));
}

export const urls = {
  docs: "https://docs.praxis.causq.com/",
  howItWorks: "https://docs.praxis.causq.com/about/how-pangolin-works",
  terms: "https://praxis.net/tos",
  privacy: "https://praxis.net/privacy",
  configureClient: "https://docs.praxis.causq.com/manage/clients/platforms"
};

/** True in the Linux build (WebKitGTK), false on Windows (WebView2). */
export const isLinux = /Linux/.test(navigator.userAgent) && !/Windows/.test(navigator.userAgent);
