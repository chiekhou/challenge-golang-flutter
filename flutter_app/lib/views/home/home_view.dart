import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import '../../providers/destination_provider.dart';
import '../../models/destination_model.dart';
import '../../widgets/app_drawer.dart';
import '../../widgets/app_loader.dart';
import 'widgets/destination_card.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import '../../providers/auth_provider.dart';

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
    DestinationProvider destinationProvider =
        Provider.of<DestinationProvider>(context);
    AuthProvider authProvider = Provider.of<AuthProvider>(context);

    List<Destination> filteredDestinations =
        destinationProvider.getFilteredDestinations(searchController.text);
    return Scaffold(
      appBar: AppBar(
        title: Text(AppLocalizations.of(context)!.title_home),
        actions: <Widget>[
          if (authProvider.isAdmin)
            IconButton(
              icon: const Icon(Icons.admin_panel_settings),
              onPressed: () {
                Navigator.pushNamed(context, '/admin/dashboard');
              },
            ),
        ],
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
                    decoration: InputDecoration(
                      hintText: AppLocalizations.of(context)!.recherche_ville,
                      prefixIcon: const Icon(
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
                child: destinationProvider.isLoading
                    ? const AppLoader()
                    : filteredDestinations.isNotEmpty
                        ? ListView.builder(
                            itemCount: filteredDestinations.length,
                            itemBuilder: (_, i) => DestinationCard(
                              destination: filteredDestinations[i],
                            ),
                          )
                        : Text(AppLocalizations.of(context)!.resultat),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
