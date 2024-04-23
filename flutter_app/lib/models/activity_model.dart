enum ActivityStatus { ongoing, done }

class Activity {
  String? id;
  String name;
  String image;
  String destination;
  double price;
  ActivityStatus status;
  LocationActivity? location;
  Activity({
    required this.name,
    required this.destination,
    required this.image,
    required this.price,
    this.id,
    this.location,
    this.status = ActivityStatus.ongoing,
  });

  Activity.fromJson(Map<String, dynamic> json)
      : id = json['_id'],
        name = json['name'],
        image = json['image'],
        destination = json['destination'],
        price = json['price'].toDouble(),
        location = LocationActivity(
          address: json['address'],
          latitude: json['latitude'],
          longitude: json['longitude'],
        ),
        status =
            json['status'] == 0 ? ActivityStatus.ongoing : ActivityStatus.done;

  Map<String, dynamic> toJson() {
    Map<String, dynamic> value = {
      'name': name,
      'image': image,
      'destination': destination,
      'price': price,
      'address': location?.address,
      'longitude': location?.longitude,
      'latitude': location?.latitude,
      'status': status == ActivityStatus.ongoing ? 0 : 1
    };
    if (id != null) {
      value['_id'] = id;
    }
    return value;
  }
}

class LocationActivity {
  String? address;
  double? longitude;
  double? latitude;
  LocationActivity({
    this.address,
    this.longitude,
    this.latitude,
  });
}
