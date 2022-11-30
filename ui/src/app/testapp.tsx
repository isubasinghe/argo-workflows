import * as React from 'react';
import {useState} from 'react';


const App = () => {
  const [tz, setTz] = useState<string>(process.env.DEFAULT_TZ);


  return (
    <div>
      <p>{tz}</p>
      <input value={tz} onChange={(el) => setTz(el.target.value)}/>
    </div>
  );
}

export default App;
