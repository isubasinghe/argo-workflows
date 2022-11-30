import * as React from 'react';
import * as ReactDOM from 'react-dom';
import App from './testapp'; 

ReactDOM.render(<App />, document.getElementById('app'));

const mdl = module as any;
if (mdl.hot) {
    mdl.hot.accept('./testapp.tsx', () => {
        const UpdatedApp = require('./testapp.tsx').App;
        ReactDOM.render(<UpdatedApp />, document.getElementById('app'));
    });
}
