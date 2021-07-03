/* eslint-disable no-new */

// Import each function separately so specific functions can be imported, rather than everything
export function IsValidUrl(url: string): boolean {
  try {
    // Use the URL class to validate the url
    new URL(url);
  } catch (_) {
    return false;
  }
  return true;
}
