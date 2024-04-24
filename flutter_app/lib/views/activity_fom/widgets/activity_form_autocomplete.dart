import 'dart:async';

import 'package:flutter/material.dart';

import '../../../apis/google_api.dart';
import '../../../models/activity_model.dart';
import '../../../models/place_model.dart';

Future showInputAutocomplete(BuildContext context) {
  return showDialog(
    context: context,
    builder: (_) => const InputAddress(),
  );
}

class InputAddress extends StatefulWidget {
  const InputAddress({super.key});

  @override
  State<InputAddress> createState() => _InputAddressState();
}

class _InputAddressState extends State<InputAddress> {
  List<Place> _places = [];
  Timer? _debounce;

  Future<void> _searchAddress(String value) async {
    try {
      if (_debounce?.isActive == true) _debounce?.cancel();
      _debounce = Timer(const Duration(seconds: 1), () async {
        if (value.isNotEmpty) {
          _places = await getAutocompleteSuggestions(value);
          setState(() {});
        }
      });
    } catch (e) {
      rethrow;
    }
  }

  Future<void> getPlaceDetails(String placeId) async {
    try {
      LocationActivity location = await getPlaceDetailsApi(placeId);
      if (mounted) {
        Navigator.pop(context, location);
      }
    } catch (e) {
      rethrow;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: <Widget>[
          Stack(
            children: <Widget>[
              TextField(
                decoration: const InputDecoration(
                  labelText: 'Rechercher',
                  prefixIcon: Icon(Icons.search),
                ),
                onChanged: _searchAddress,
              ),
              Positioned(
                top: 5,
                right: 3,
                child: IconButton(
                  icon: const Icon(Icons.clear),
                  onPressed: () => Navigator.pop(context, null),
                ),
              )
            ],
          ),
          Expanded(
            child: ListView.builder(
              itemCount: _places.length,
              itemBuilder: (_, i) {
                var place = _places[i];
                return ListTile(
                  leading: const Icon(Icons.place),
                  title: Text(place.description),
                  onTap: () => getPlaceDetails(place.placeId),
                );
              },
            ),
          )
        ],
      ),
    );
  }
}
