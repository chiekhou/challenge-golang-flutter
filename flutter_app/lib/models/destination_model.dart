import 'activity_model.dart';

class Destination {
  String? id;
  String image;
  String name;
  List<Activity> activities;
  Destination({
    required this.id,
    required this.image,
    required this.name,
    required this.activities,
  });

  Destination.fromJson(Map<String, dynamic> json)
      : id = json['_id'],
        image = json['image'],
        name = json['name'],
        activities = (json['activities'] as List)
            .map((activityJson) => Activity.fromJson(activityJson))
            .toList();
}
