enum ActivityStatus { ongoing, done }

class Activity {
  int id;
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
    required this.id,
    this.location,
    this.status = ActivityStatus.ongoing,
  });

  String toString() {
    return 'Activity{id: $id, name: $name, destination: $destination, price: $price}';
  }

  Activity.fromJson(Map<String, dynamic> json)
      : id = json['id'],
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
      'id': id,
      'name': name,
      'image': image,
      'destination': destination,
      'price': price,
      'address': location?.address,
      'longitude': location?.longitude,
      'latitude': location?.latitude,
      'status': status == ActivityStatus.ongoing ? 0 : 1
    };
    value['id'] = id;
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
