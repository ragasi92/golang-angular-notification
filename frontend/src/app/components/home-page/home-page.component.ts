import { Component, OnInit } from '@angular/core';
import { Subscription, forkJoin, of, Subject, timer, from } from 'rxjs';
import { catchError, takeUntil, concatMap, tap } from 'rxjs/operators';
import { NotificationsService } from '../../core/notifications.service';
import { AppNotification } from '../../model/notification';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {
  title = 'client';
  private subscription: Subscription;
  private unReadNotifications: AppNotification[];
  public numberOfNotifications: number;
  public notifications: AppNotification[];
  public selectedNotification: AppNotification;
  private componentDestroyed = new Subject();
  constructor(
    private notificationService: NotificationsService
  ) {
    this.numberOfNotifications = 0;
  }

  ngOnInit(): void {
    this.loadNotifications()
  }

  private loadNotifications(): void {
    timer(0, 5000)
      .pipe(takeUntil(this.componentDestroyed))
      .pipe(concatMap(() => from(this.notificationService.getNotifications()))).subscribe((response: AppNotification[]) => {
        this.processResponse(response);
      });
  }

  public markAsRead(notification: AppNotification): void {
    const id = notification.id;
    this.subscription = this.notificationService.updateNotification(id)
      .pipe(
        concatMap(() => {
          return this.notificationService.getNotifications()
        })
      ).subscribe(
        (response: AppNotification[]) => {
          this.processResponse(response)
        }
      );
  }

  private processResponse(response: AppNotification[]) {
    this.notifications = response;
    this.unReadNotifications = this.notifications.filter((notification) => {
      return notification.status == "UNREAD"
    })
    this.numberOfNotifications = this.unReadNotifications.length
  }

  onRowSelect(event: any) {
    console.log(event)
  }

  ngOnDestroy(): void {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
    this.componentDestroyed.unsubscribe();
  }

}
