import React from 'react';
import FBILogo from '../res/fbi.png';

import './compromised.scss';

const Compromised = () => (
  <div className="compromised-container">
    <h1>THIS DOMAIN HAS BEEN SEIZED</h1>
    <p>
      This domain has been seized by the Federal Bureau of Investigation pursuant to a seizure
      warrant issued by the United States District Court for the District of
      Columbia under the authority of 18 U.S.C. &sect;&sect; 981, 982, inter alia, as part of
      coordinated law enforcement action by:
    </p>
    <img className="logo" src={FBILogo} alt="FBI logo" />
  </div>
);

export default Compromised;
