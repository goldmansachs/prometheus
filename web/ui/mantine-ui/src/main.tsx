import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import store from "./state/store";
import { Provider } from "react-redux";
import "./fonts/codicon.ttf";
import "./promql.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <App />
    </Provider>
  </React.StrictMode>
);
