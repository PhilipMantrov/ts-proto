import { Component, OnInit } from '@angular/core';
import { DashStateClientImpl, GrpcWebImpl } from '../example';
import { grpc } from '@improbable-eng/grpc-web';

@Component({
  selector: 'app-root',
  template: ``,
  styles: []
})
export class AppComponent implements OnInit {
  title = 'angular-frontend';

  constructor() {}

  private readonly rpc: GrpcWebImpl = new GrpcWebImpl("http://localhost:8080", {
    transport: grpc.WebsocketTransport(),
    debug: false,
    metadata: new grpc.Metadata({ SomeHeader: "bar" }),
  });

  private readonly dashStateClientImpl: DashStateClientImpl = new DashStateClientImpl(this.rpc);

  public ngOnInit() {
    this.dashStateClientImpl.UserSettings({}).subscribe(next => {
      this.dashStateClientImpl.SendUserSetting(next).subscribe();
    });
  }
}
