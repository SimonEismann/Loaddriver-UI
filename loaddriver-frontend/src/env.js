const apiRoot = () => {
  if (window.env && window.env.API_ROOT) {
    return window.env.API_ROOT;
  } else {
    return 'http://localhost';
  }
};

const consoleUri = () => {
  if (window.env && window.env.CONSOLE_URI) {
    return window.env.CONSOLE_URI;
  } else {
    return 'ws://localhost/jobs/current/output';
  }
};


export const API_ROOT = apiRoot();
export const CONSOLE_URI = consoleUri();