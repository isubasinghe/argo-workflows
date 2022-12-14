import {Page} from 'argo-ui';
import {BadMagic} from 'badmagic';
import * as React from 'react';
import {useEffect, useState} from 'react';
import 'swagger-ui-react/swagger-ui.css';
import {uiUrl} from '../../shared/base';
import {useCollectEvent} from '../../shared/components/use-collect-event';
require('../style.scss');
export const ApiDocs = () => {
    useCollectEvent('openedApiDocs');
    const [swagger, setSwagger] = useState(null);

    useEffect(() => {
        fetch(uiUrl('assets/openapi-spec/swagger.json'))
            .then(res => res.json())
            .then(json => {
                console.log(json);
                setSwagger(json);
            })
            .catch(err => console.log(err));
    }, []);
    return (
        <Page title='Swagger'>
            <div className='argo-container'>
                {swagger && (
                    <BadMagic
                        workspaces={[{id: 'main', name: 'main', routes: [], openapi: JSON.stringify(swagger), config: {baseUrl: uiUrl('assets/openapi-spec/swagger.json')}}]}
                    />
                )}
            </div>
        </Page>
    );
};
