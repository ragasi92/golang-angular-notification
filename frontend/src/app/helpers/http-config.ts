import { HttpHeaders, HttpParams } from '@angular/common/http';

const headerJson = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
    'Access-Control-Allow-Origin': 'http://localhost:12345'
};

const headerFormData = {
};

export const contentHeaders = new HttpHeaders(headerJson);

export const contentHeadersFormData = new HttpHeaders(headerFormData);

export class AppHttpParams {
    params: HttpParams;

    constructor() {
        this.params = new HttpParams();
    }

    public set(param: string, value: string) {
        this.params = this.params.set(param, value);
    }

    public getParams(): HttpParams {
        return this.params;
    }

    /**
     * This wil hide the spinner from the ui for an API call
     * usage
     *    const params = new AytHttpParams();
     *    params.disableSpinner();
     *    and pass the params into API service method
     * */
    public disableSpinner() {
        this.params = this.params.set('spinner', 'false');
    }
}
