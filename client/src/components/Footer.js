import React from "react";

export const Footer = () => {
  return (
    <footer className="footer" style={{ left: "0" }}>
      <div className="container">
        <span className="text-muted">&copy; 2019-2020 CAP THAT PIC</span>
        <span className="text-muted ml-5">
          <a href="https://github.com/CoolDogee/cap-that-pic" target="_blank">
            Contribute
          </a>
        </span>
      </div>
    </footer>
  );
};

export default Footer;
