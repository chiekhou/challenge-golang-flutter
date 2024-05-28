import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../providers/destination_provider.dart';
import '../../models/destination_model.dart';
import '../../widgets/app_drawer.dart';
import '../../widgets/app_loader.dart';
import 'widgets/destination_card.dart';

class HomeView extends StatefulWidget {
  static const String routeName = '/home';

  const HomeView({super.key});

  @override
  State<HomeView> createState() {
    return _HomeState();
  }
}

class _HomeState extends State<HomeView> {
  TextEditingController searchController = TextEditingController();

  @override
  void initState() {
    super.initState();
    searchController.addListener(() {
      setState(() {});
    });
  }

  @override
  void dispose() {
    searchController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    DestinationProvider cityProvider =
        Provider.of<DestinationProvider>(context);
    List<Destination> filteredDestinations =
        cityProvider.getFilteredCities(searchController.text);
    return Scaffold(
      appBar: AppBar(
        title: const Text('esgivoyage'),
      ),
      drawer: const AppDrawer(),
      body: Column(
        children: <Widget>[
          Container(
            margin: const EdgeInsets.only(top: 10),
            padding: const EdgeInsets.symmetric(horizontal: 14),
            child: Row(
              children: <Widget>[
                Expanded(
                  child: TextField(
                    controller: searchController,
                    decoration: const InputDecoration(
                      hintText: 'Rechercher une ville',
                      prefixIcon: Icon(
                        Icons.search,
                      ),
                    ),
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.clear),
                  onPressed: () => setState(() => searchController.clear()),
                )
              ],
            ),
          ),
          Expanded(
            child: Container(
              padding: const EdgeInsets.all(10),
              child: RefreshIndicator(
                displacement: 100.0,
                onRefresh:
                    Provider.of<DestinationProvider>(context, listen: false)
                        .fetchData,
                child: cityProvider.isLoading
                    ? const AppLoader()
                    : filteredDestinations.isNotEmpty
                        ? ListView.builder(
                            itemCount: filteredDestinations.length,
                            itemBuilder: (_, i) => DestinationCard(
                              destination: filteredDestinations[i],
                            ),
                          )
                        : const Text('Aucun résultat'),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
