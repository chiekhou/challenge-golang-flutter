import 'activity_model.dart';

class Voyage {
  int? id;
  String destination;
  List<Activity> activities;
  DateTime? date;
  Voyage({
    required this.destination,
    required this.activities,
    this.date,
    this.id,
  });

  @override
  String toString() {
    return 'Voyage{id: $id, destination: $destination, date: $date, activities: $activities}';
  }

  factory Voyage.fromJson(Map<String, dynamic> json) {
    var activitiesFromJson = json['activities'] as List;
    List<Activity> activityList = activitiesFromJson
        .map((activityJson) =>
            Activity.fromJson(activityJson as Map<String, dynamic>))
        .toList();

    return Voyage(
      id: json['id'],
      activities: activityList,
      date: DateTime.parse(json['date']),
      destination: json['destination'],
    );
  }
  Map<String, dynamic> toJson() {
    final dateFormatted = date?.toIso8601String();
    return {
      'id': id,
      'activities': activities.map((activity) => activity.toJson()).toList(),
      'date': dateFormatted,
      'destination': destination,
    };
  }
}
