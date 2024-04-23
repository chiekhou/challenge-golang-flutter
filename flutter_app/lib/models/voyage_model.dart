import 'activity_model.dart';

class Voyage {
  String? id;
  String destination;
  List<Activity> activities;
  DateTime? date;
  Voyage({
    this.id,
    required this.destination,
    required this.activities,
    this.date,
  });

  Voyage.fromJson(Map<String, dynamic> json)
      : id = json['_id'],
        destination = json['destination'],
        date = DateTime.parse(json['date']),
        activities = (json['activities'] as List)
            .map(
              (activityJson) => Activity.fromJson(activityJson),
            )
            .toList();

  Map<String, dynamic> toJson() {
    if (id != null) {
      return {
        '_id': id,
        'destination': destination,
        'date': date!.toIso8601String(),
        'activities': activities
            .map(
              (activity) => activity.toJson(),
            )
            .toList()
      };
    } else {
      return {
        'destination': destination,
        'date': date!.toIso8601String(),
        'activities': activities
            .map(
              (activity) => activity.toJson(),
            )
            .toList(),
      };
    }
  }
}
