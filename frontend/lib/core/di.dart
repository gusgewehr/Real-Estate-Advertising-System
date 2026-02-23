import 'package:frontend/core/api_client.dart';
import 'package:frontend/data/datasources/exchange_rate.dart';
import 'package:frontend/data/datasources/real_estate.dart';
import 'package:frontend/domain/usecase/real_estate.dart';
import 'package:get_it/get_it.dart';

import '../data/datasources/address.dart';
import '../data/repository/address.dart';
import '../data/repository/exchange_rate.dart';
import '../data/repository/real_estate.dart';
import '../domain/repository/address.dart';
import '../domain/repository/exchange_rate.dart';
import '../domain/repository/real_estate.dart';
import '../domain/usecase/address.dart';
import '../domain/usecase/exchange_rate.dart';
import '../presentation/controller/address.dart';
import '../presentation/controller/exchange_rate.dart';
import '../presentation/controller/real_estate.dart';

final sl = GetIt.instance;

void setupDependencies() {

  sl.registerLazySingleton<ApiClient>(() => ApiClient());

  // Real Estate
  sl.registerLazySingleton<RealEstateDataSource>(
      () => RealEstateDataSourceImpl(apiClient: sl()),
  );
  sl.registerLazySingleton<RealEstateRepository>(
      () => RealEstateRepositoryImpl(realEstateDataSource: sl()),
  );
  sl.registerLazySingleton<GetRealEstateUseCase>(
      () => GetRealEstateUseCase(realEstateRepository: sl()),
  );
  sl.registerLazySingleton<UploadImageUseCase>(
      () => UploadImageUseCase(realEstateRepository: sl()),
  );
  sl.registerLazySingleton<CreateRealEstateUseCase>(
      () => CreateRealEstateUseCase(realEstateRepository: sl()),
  );
  sl.registerFactory<RealEstateController>(
      () => RealEstateController(getRealEstateUseCase: sl(), uploadImageUseCase: sl(), createRealEstateUseCase: sl()),
  );


  // Exchange Rate
  sl.registerLazySingleton<ExchangeRateDataSource>(
      () => ExchangeRateDataSourceImpl(apiClient: sl()),
  );
  sl.registerLazySingleton<ExchangeRateRepository>(
      () => ExchangeRateRepositoryImpl(exchangeRateDataSource: sl()),
  );
  sl.registerLazySingleton<GetLatestExchangeRateUseCase>(
      () => GetLatestExchangeRateUseCase(exchangeRateRepository: sl()),
  );
  sl.registerLazySingleton<AddExchangeRateUseCase>(
      () => AddExchangeRateUseCase(exchangeRateRepository: sl()),
  );
  sl.registerFactory<ExchangeRateController>(
      () => ExchangeRateController(getLatestExchangeRateUseCase: sl(), addExchangeRateUseCase: sl()),
  );


  // Address
  sl.registerLazySingleton<AddressDataSource>(
      () => AddressDataSourceImpl(apiClient: sl()),
  );
  sl.registerLazySingleton<AddressRepository>(
      () => AddressRepositoryImpl(addressDataSource: sl()),
  );
  sl.registerLazySingleton<GetAddressUseCase>(
      () => GetAddressUseCase(addressRepository: sl()),
  );
  sl.registerFactory<AddressController>(
      () => AddressController(getAddressUseCase: sl()),
  );





}

