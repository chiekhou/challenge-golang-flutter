import 'package:flutter/material.dart';
import 'package:flutter_app/models/hotel_model.dart';
import 'package:flutter_app/views/destination/widgets/hotel_card.dart';


class HotelList extends StatelessWidget {
  final List<Hotel> hotels;
  final List<Hotel> selectedHotels;
  final Function toggleHotel;

  const HotelList({
    super.key,
    required this.hotels,
    required this.selectedHotels,
    required this.toggleHotel,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      mainAxisSize: MainAxisSize.min,
      children: [
        Container(
          padding: EdgeInsets.all(8.0),
          color: Colors.grey[300],
          child: Text(
            'Liste des Hôtels',
            style: TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
          ),
        ),

        Flexible(
          fit: FlexFit.loose,
          child: GridView.builder(
            shrinkWrap: true,
            physics: NeverScrollableScrollPhysics(),
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 2,
              crossAxisSpacing: 10,
              mainAxisSpacing: 10,
              childAspectRatio: 3 / 2,
            ),
            itemCount: hotels.length,
            itemBuilder: (context, index) {
              final hotel = hotels[index];
              return HotelCard(
                hotel: hotel,
                isSelected: selectedHotels.contains(hotel),
                toggleHotel: () => toggleHotel(hotel),
              );
            },
          ),
        ),
      ],
    );
  }
}



