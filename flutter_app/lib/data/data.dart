import '../models/voyage_model.dart';
import '../models/activity_model.dart';

List<Voyage> voyages = [
  Voyage(
    activities: [
      Activity(
        image: 'assets/images/activities/louvre.jpeg',
        name: 'Louvre',
        id: 'a1',
        destination: 'Paris',
        price: 12.00,
      ),
      Activity(
        image: 'assets/images/activities/chaumont.jpeg',
        name: 'Chaumont',
        id: 'a2',
        destination: 'Paris',
        price: 0.00,
      ),
      Activity(
        image: 'assets/images/activities/dame.jpeg',
        name: 'Notre-Dame',
        id: 'a3',
        destination: 'Paris',
        price: 0.00,
      ),
    ],
    destination: 'Paris',
    date: DateTime.now().add(const Duration(days: 1)),
  ),
  Voyage(
      activities: [],
      destination: 'Lyon',
      date: DateTime.now().add(const Duration(days: 2))),
  Voyage(
      activities: [],
      destination: 'Nice',
      date: DateTime.now().subtract(const Duration(days: 1))),
];
