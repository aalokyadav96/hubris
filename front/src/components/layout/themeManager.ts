import { getActiveDomainMetadata } from "../../config/domainFeatures.js";

export const themes = [
  "light",
  "dark",
  "dimmed",
  "solarized",
  "nord",
  "midnight",
  "emerald",
  "dracula",
  "catppuccin",
  "gruvbox",
  "tokyo-night",
  "cyberpunk",
  "latte",
  "rose-pine",
  "high-contrast"
] as const;

export type Theme = typeof themes[number];

let currentThemeIndex = 0;

function clearDomainThemeClasses(root: HTMLElement): void {
  const domainThemeClasses = Array.from(root.classList).filter((className) => className.startsWith("theme-"));
  domainThemeClasses.forEach((className) => root.classList.remove(className));
}

/**
 * Applies a user theme mode and coordinates domain branding classes
 * 
 * @param theme - The user theme key (e.g., 'dark', 'light')
 * @param save - Whether to persist to localStorage
 */
export function applyTheme(theme: Theme, save: boolean = true): void {
  // Use explicit type assertion to satisfy TypeScript's readonly array check
  if (!(themes as readonly string[]).includes(theme)) return;

  const root = document.documentElement;

  // 1. Set user color scheme attribute using bracket notation for index signature safety
  root.dataset['theme'] = theme;
  currentThemeIndex = themes.indexOf(theme);

  // 2. Clear any stale domain branding classes before adding the active one
  clearDomainThemeClasses(root);
  const domainMeta = getActiveDomainMetadata() as { theme?: string } | undefined;
  const domainTheme = domainMeta?.theme;
  if (domainTheme) {
    root.classList.add(domainTheme);
  }

  // 3. Persist to localStorage if requested
  if (save) {
    try {
      localStorage.setItem("theme", theme);
    } catch {
      // Ignore storage write failures.
    }
  }
}

/**
 * Loads stored theme, falls back to OS preference, and applies domain theme
 */
export function loadTheme(): void {
  try {
    const saved = localStorage.getItem("theme");

    if (saved && (themes as readonly string[]).includes(saved)) {
      applyTheme(saved as Theme, false);
      return;
    }
  } catch {
    // Ignore storage access failures.
  }

  // Fall back to OS preference if no theme saved
  const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const defaultTheme: Theme = prefersDark ? "dark" : "light";
  applyTheme(defaultTheme, false);
}

/**
 * Cycles through available themes
 */
export function toggleTheme(): void {
  currentThemeIndex = (currentThemeIndex + 1) % themes.length;
  // Fall back to "light" or ensure it's defined 👇
  const theme = themes[currentThemeIndex] ?? "light";
  applyTheme(theme, true);
}

/**
 * Reacts to system dark/light mode changes if the user hasn't set an explicit override
 */
export function listenForSystemThemeChanges(): void {
  window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", (e) => {
    try {
      if (!localStorage.getItem("theme")) {
        applyTheme(e.matches ? "dark" : "light", false);
      }
    } catch {
      // Ignore storage access failures.
    }
  });
}