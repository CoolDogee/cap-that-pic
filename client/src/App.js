import React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import "./App.css";

import routes from "./routes";
import withTracker from "./withTracker";
import 'bootstrap/dist/css/bootstrap.min.css';

export default () => (
  <Router basename={process.env.REACT_APP_BASENAME || ""}>
    <div className="container mt-4">
      {routes.map((route, index) => {
        return (
          <Route
            key={index}
            path={route.path}
            exact={route.exact}
            component={withTracker(props => {
              return (
                <route.layout {...props}>
                  <route.component {...props} />
                </route.layout>
              );
            })}
          />
        );
      })}
    </div>
  </Router>
);
