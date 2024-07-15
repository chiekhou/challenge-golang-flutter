import 'package:awesome_notifications/awesome_notifications.dart';

class NotificationsController {
  @pragma("vm:entry-point")
  static Future<void> onNotificationMethod(
      ReceivedNotification receivedNotification)async{}

  @pragma("vm:entry-point")
  static Future<void> onNotificationDisplayedMethod(
      ReceivedNotification receivedNotification)async{}

  static Future<void> onDismissActionReceivedethod(
      ReceivedNotification receivedNotification)async{}

  static Future<void> onActionReceivedMethod(
      ReceivedNotification receivedNotification)async{}

  static Future<void> onNotificationCreatedMethod(
      ReceivedNotification receivedNotification)async{}
}