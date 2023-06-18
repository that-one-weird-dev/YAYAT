
const port = location.port === "3000" ? 3001 : location.port;
export const ENDPOINT = `http://${location.hostname}:${port}` as const;
