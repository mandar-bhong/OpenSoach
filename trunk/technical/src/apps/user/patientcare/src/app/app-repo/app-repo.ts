export class AppRepo {

    private static singleton: AppRepo;


    API_SPL_BASE_URL: string;
    API_APP_BASE_URL: string;


    private constructor() {
        this.API_SPL_BASE_URL = "http://172.105.232.148/api";
        this.API_APP_BASE_URL = "http://172.105.232.148:91/api";
    }


    public static getInstance(): AppRepo {
        if (!AppRepo.singleton) {
            AppRepo.singleton = new AppRepo();
        }

        return AppRepo.singleton;
    }



}