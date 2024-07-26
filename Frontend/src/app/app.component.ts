declare var google: any;
import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private httpClient: HttpClient) {}

  ngOnInit() {
    google.accounts.id.initialize({
      client_id: "1034237727696-umachtrijarh84r34etrf3asll37ubbk.apps.googleusercontent.com",
      callback: (response: any) => this.handleCredentialResponse(response)
    });
    google.accounts.id.renderButton(document.getElementById("google-btn"), {
      theme: 'filled_blue', size: 'large', shape: 'rectangle', width: 350
    });
  }

  handleCredentialResponse(response: any) {
    console.log(response);
    this.httpClient.post('http://localhost:8000/api/google-login', { token: response.credential })
      .subscribe(result => {
        console.log('Backend response', result);
      });
  }
}
